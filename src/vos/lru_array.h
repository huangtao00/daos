/**
 * (C) Copyright 2020 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */
/**
 * Generic struct for allocating LRU entries in an array
 * common/lru_array.h
 *
 * Author: Jeff Olivier <jeffrey.v.olivier@intel.com>
 */

#ifndef __LRU_ARRAY__
#define __LRU_ARRAY__

#include <daos/common.h>

struct lru_callbacks {
	/** Called when an entry is going to be evicted from cache */
	void	(*lru_on_evict)(void *entry, uint32_t idx, void *arg);
	/** Called on initialization of an entry */
	void	(*lru_on_init)(void *entry, uint32_t idx, void *arg);
	/** Called on finalization of an entry */
	void	(*lru_on_fini)(void *entry, uint32_t idx, void *arg);
};

struct lru_entry {
	/** The pointer to the index is unique identifier for the entry */
	uint64_t	 le_key;
	/** Pointer to this entry */
	void		*le_payload;
	/** Next index in LRU array */
	uint32_t	 le_next_idx;
	/** Previous index in LRU array */
	uint32_t	 le_prev_idx;
};

struct lru_sub {
	/** Index of LRU */
	uint32_t		 ls_lru;
	/** Index of first free entry */
	uint32_t		 ls_free;
	/** Index of this entry in the array */
	uint32_t		 ls_array_idx;
	/** Padding */
	uint32_t		 ls_pad;
	/** Link in the array free/unused list.  If the subarray has no free
	 *  entries, it is removed from either list so this field is unused.
	 */
	d_list_t		 ls_link;
	/** Allocated payload entries */
	void			*ls_payload;
	/** Entries in the array */
	struct lru_entry	*ls_table;
};

#define LRU_NO_IDX	0xffffffff

enum {
	/** Free'd entries are added to tail of free list */
	LRU_FLAG_EVICT_MANUAL		= 1,
	/** Free'd entries are added to tail of free list */
	LRU_FLAG_REUSE_UNIQUE		= 2,
};

struct lru_array {
	/** Number of indices */
	uint32_t		 la_count;
	/** record size */
	uint16_t		 la_record_size;
	/** eviction count */
	uint16_t		 la_evicting;
	/** Array flags */
	uint32_t		 la_flags;
	/** Second level mask */
	uint32_t		 la_array_mask;
	/** Second level bit shift */
	uint32_t		 la_array_shift;
	/** First level mask */
	uint32_t		 la_idx_mask;
	/** Subarrays with free entries */
	d_list_t		 la_free_sub;
	/** Unallocated subarrays */
	d_list_t		 la_unused_sub;
	/** Callbacks for implementation */
	struct lru_callbacks	 la_cbs;
	/** User callback argument passed on init */
	void			*la_arg;
	/** Allocated subarrays */
	struct lru_sub		 la_sub[0];
};

/** Internal API: Evict the LRU, move it to MRU, invoke eviction callback,
 *  and return the index
 */
void
lrua_find_free(struct lru_array *array, struct lru_entry **entry,
	       uint32_t *idx, uint64_t key);

/** Internal API: Remove an entry from the lru list */
static inline void
lrua_remove_entry(struct lru_sub *sub, uint32_t *head, struct lru_entry *entry,
		  uint32_t idx)
{
	struct lru_entry	*entries = &sub->ls_table[0];
	struct lru_entry	*prev = &entries[entry->le_prev_idx];
	struct lru_entry	*next = &entries[entry->le_next_idx];

	/** Last entry in the list */
	if (prev == entry) {
		*head = LRU_NO_IDX;
		return;
	}

	prev->le_next_idx = entry->le_next_idx;
	next->le_prev_idx = entry->le_prev_idx;

	if (idx == *head)
		*head = entry->le_next_idx;
}

/** Internal API: Insert an entry in the lru list */
static inline void
lrua_insert(struct lru_sub *sub, uint32_t *head, struct lru_entry *entry,
	    uint32_t idx, bool append)
{
	struct lru_entry	*entries = &sub->ls_table[0];
	struct lru_entry	*prev;
	struct lru_entry	*next;
	uint32_t		 tail;

	if (*head == LRU_NO_IDX) {
		*head = entry->le_prev_idx = entry->le_next_idx = idx;
		return;
	}

	next = &entries[*head];
	tail = next->le_prev_idx;
	prev = &entries[tail];
	next->le_prev_idx = idx;
	prev->le_next_idx = idx;
	entry->le_prev_idx = tail;
	entry->le_next_idx = *head;

	if (append)
		return;

	*head = idx;
}

/** Internal API: Make the entry the mru */
static inline void
lrua_move_to_mru(struct lru_sub *sub, struct lru_entry *entry, uint32_t idx)
{
	if (entry->le_next_idx == sub->ls_lru) {
		/** Already the mru */
		return;
	}

	if (sub->ls_lru == idx) {
		/** Ordering doesn't change in circular list so just update
		 *  the lru and mru idx
		 */
		sub->ls_lru = entry->le_next_idx;
		return;
	}

	/** First remove */
	lrua_remove_entry(sub, &sub->ls_lru, entry, idx);

	/** Insert at mru */
	lrua_insert(sub, &sub->ls_lru, entry, idx, true);
}

/** Internal API to lookup entry from index */
static inline struct lru_entry *
lrua_lookup_idx(struct lru_array *array, uint32_t idx, uint64_t key)
{
	struct lru_entry	*entry;
	struct lru_sub		*sub;
	uint32_t		 sub_idx;
	uint32_t		 ent_idx;

	if (idx >= array->la_count)
		return NULL;

	sub_idx = (idx & array->la_array_mask) >> array->la_array_shift;
	ent_idx = idx & array->la_idx_mask;

	sub = &array->la_sub[sub_idx];
	entry = &sub->ls_table[ent_idx];
	if (entry->le_key == key) {
		if (!array->la_evicting) {
			/** Only make mru if we are not evicting it */
			lrua_move_to_mru(sub, entry, ent_idx);
		}
		return entry;
	}

	return NULL;
}

/** Lookup an entry in the lru array with alternative key.
 *
 * \param	array[in]	The lru array
 * \param	idx[in]		The index of the entry
 * \param	idx[in]		Unique identifier
 * \param	entryp[in,out]	Valid only if function returns true.
 *
 * \return true if the entry is in the array and set \p entryp accordingly
 */
static inline bool
lrua_lookupx(struct lru_array *array, uint32_t idx, uint64_t key,
	     void **entryp)
{
	struct lru_entry	*entry;

	D_ASSERT(array != NULL);
	D_ASSERT(key != 0);

	*entryp = NULL;

	entry = lrua_lookup_idx(array, idx, key);
	if (entry == NULL)
		return false;

	*entryp = entry->le_payload;
	return true;
}

/** Lookup an entry in the lru array.
 *
 * \param	array[in]	The lru array
 * \param	idx[in,out]	Address of the record index.
 * \param	entryp[in,out]	Valid only if function returns true.
 *
 * \return true if the entry is in the array and set \p entryp accordingly
 */
static inline bool
lrua_lookup(struct lru_array *array, const uint32_t *idx,
	    void **entryp)
{
	return lrua_lookupx(array, *idx, (uint64_t)idx, entryp);
}

/** Allocate a new entry lru array with alternate key specifier.
 *  This should only be called if lookup would return false.  This will
 *  modify idx.  If called within a transaction and the value needs to
 *  persist, the old value should be logged before calling this function.
 *
 * \param	array[in]	The LRU array
 * \param	idx[in,out]	Index address in, allocated index out
 * \param	key[in]		Unique identifier of entry
 *
 * \return	Returns a pointer to the entry.  It can return NULL if
 *		manual eviction flag is set and either there are no available
 *		entries or no an allocation failed.
 */
static inline void *
lrua_allocx(struct lru_array *array, uint32_t *idx, uint64_t key)
{
	struct lru_entry	*new_entry;

	D_ASSERT(array != NULL);
	D_ASSERT(key != 0);

	lrua_find_free(array, &new_entry, idx, key);

	if (new_entry == NULL)
		return NULL;

	return new_entry->le_payload;
}

/** Allocate a new entry lru array.   This should only be called if lookup
 *  would return false.  This will modify idx.  If called within a
 *  transaction and the value needs to persist, the old value should be
 *  logged before calling this function.
 *
 * \param	array[in]	The LRU array
 * \param	idx[in,out]	Address of the entry index.
 *
 * \return	Returns a pointer to the entry.  It can return NULL if
 *		manual eviction flag is set and either there are no available
 *		entries or no an allocation failed.
 */
static inline void *
lrua_alloc(struct lru_array *array, uint32_t *idx)
{
	return lrua_allocx(array, idx, (uint64_t)idx);
}

/** Allocate an entry in place.  Used for recreating an old array.
 *
 * \param	array[in]	The LRU array
 * \param	idx[in]		Index of entry.
 * \param	key[in]		Address of the entry index.
 *
 * \return	Returns a pointer to the entry or NULL on error
 */
static inline void *
lrua_allocx_inplace(struct lru_array *array, uint32_t idx, uint64_t key)
{
	struct lru_entry	*entry;
	struct lru_sub		*sub;
	uint32_t		 sub_idx;
	uint32_t		 ent_idx;

	D_ASSERT(array != NULL);
	D_ASSERT(key != 0);

	if (idx >= array->la_count) {
		D_ERROR("Index %d is out of range\n", idx);
		return NULL;
	}

	sub_idx = (idx & array->la_array_mask) >> array->la_array_shift;
	ent_idx = idx & array->la_idx_mask;
	sub = &array->la_sub[sub_idx];
	entry = &sub->ls_table[ent_idx];
	if (entry->le_key != key && entry->le_key != 0) {
		D_ERROR("Cannot allocated idx %d in place\n", idx);
		return NULL;
	}

	entry->le_key = key;
	lrua_move_to_mru(sub, entry, ent_idx);

	return entry->le_payload;

}

/** If an entry is still in the array, evict it and invoke eviction callback.
 *  Move the evicted entry to the LRU and mark it as already evicted.
 *
 * \param	array[in]	Address of the LRU array.
 * \param	idx[in]		Index of the entry
 * \param	key[in]		Unique identifier
 */
void
lrua_evictx(struct lru_array *array, uint32_t idx, uint64_t key);

/** If an entry is still in the array, evict it and invoke eviction callback.
 *  Move the evicted entry to the LRU and mark it as already evicted.
 *
 * \param	array[in]	Address of the LRU array.
 * \param	idx[in]		Address of the entry index.
 */
static inline void
lrua_evict(struct lru_array *array, uint32_t *idx)
{
	lrua_evictx(array, *idx, (uint64_t)idx);
}

/** Allocate an LRU array
 *
 * \param	array[in,out]	Pointer to LRU array
 * \param	nr_ent[in]	Number of records in array
 * \param	nr_arrays[in]	Number of 2nd level arrays.   If it is not 1,
 *				manual eviction is implied.
 * \param	rec_size[in]	Size of each record
 * \param	cbs[in]		Optional callbacks
 * \param	arg[in]		Optional argument passed to all callbacks
 *
 * \return	-DER_NOMEM	Not enough memory available
 *		0		Success
 */
int
lrua_array_alloc(struct lru_array **array, uint32_t nr_ent, uint32_t nr_arrays,
		 uint16_t record_size, uint32_t flags,
		 const struct lru_callbacks *cbs, void *arg);

/** Free an LRU array
 *
 * \param	array[in]	Pointer to LRU array
 */
void
lrua_array_free(struct lru_array *array);

#endif /* __LRU_ARRAY__ */
