#!/usr/bin/python
"""
  (C) Copyright 2020 floatel Corporation.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

  GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
  The Government's rights to use, modify, reproduce, release, perform, display,
  or disclose this software are subject to the terms of the Apache License as
  provided in Contract No. B609815.
  Any reproduction of computer software, computer software documentation, or
  portions thereof marked with this legend must also reproduce the markings.
"""

import time
import threading
import uuid
from itertools import product

from apricot import TestWithServers
from write_host_file import write_host_file
from test_utils_pool import TestPool
from test_utils_container import TestContainer
from ior_utils import IorCommand
from job_manager_utils import Mpirun
from command_utils_base import CommandFailure
from mpio_utils import MpioUtils

try:
    # python 3.x
    import queue as queue
except ImportError:
    # python 2.7
    import Queue as queue


class NvmePoolCapacity(TestWithServers):
    # pylfloat: disable=too-many-ancestors
    """Test class Description: Verify NOSPC
    condition is reported when accessing data beyond
    pool size.

    :avocado: recursive
    """

    def setUp(self):
        """Set up for test case."""
        super(NvmePoolCapacity, self).setUp()

        self.ior_flags = self.params.get("ior_flags", '/run/ior/iorflags/*')
        self.ior_apis = self.params.get("ior_api", '/run/ior/iorflags/*')
        self.ior_test_sequence = self.params.get("ior_test_sequence",
                                                 '/run/ior/iorflags/*')
        self.ior_daos_oclass = self.params.get("obj_class",
                                               '/run/ior/iorflags/*')
        # Recreate the client hostfile without slots defined
        self.hostfile_clients = write_host_file(
            self.hostlist_clients, self.workdir, None)
        self.pool = None
        self.out_queue = queue.Queue()

    def ior_thread(self, pool, oclass, api, test, flags, results):
        """Start threads and wait until all threads are finished.
        Args:
            pool: pool handle
            oclass: IOR object class
            API : IOR API
            test : IOR test sequence
            flags : IOR flags
            results (queue): queue for returning thread results

        Returns:
            None
        """
        processes = self.params.get("slots", "/run/ior/clientslots/*")
        container_info = {}
        mpio_util = MpioUtils()
        if mpio_util.mpich_installed(self.hostlist_clients) is False:
            self.fail("Exiting Test: Mpich not installed")
        self.pool = pool
        # Define the arguments for the ior_runner_thread method
        ior_cmd = IorCommand()
        ior_cmd.get_params(self)
        ior_cmd.set_daos_params(self.server_group, self.pool)
        ior_cmd.daos_oclass.update(oclass)
        ior_cmd.api.update(api)
        ior_cmd.transfer_size.update(test[2])
        ior_cmd.block_size.update(test[3])
        ior_cmd.flags.update(flags)

        container_info["{}{}{}"
                       .format(oclass,
                               api,
                               test[2])] = str(uuid.uuid4())

        # Define the job manager for the IOR command
        manager = Mpirun(ior_cmd, mpitype="mpich")
        manager.job.daos_cont.update(container_info
                                     ["{}{}{}".format(oclass,
                                                      api,
                                                      test[2])])
        env = ior_cmd.get_default_env(str(manager))
        manager.setup_command(env, self.hostfile_clients,
                              processes)

        # run IOR Command
        try:
            manager.run()
        except CommandFailure as _error:
            results.put("FAIL")

    def test_create_delete(self, num_pool=2, num_cont=5, total_count=100,
                           scm_size=100000000000, nvme_size=300000000000):
        """
        Test Description:
            This method is used to create/delete pools
            for a long run. It verifies the NVME free space
            during this process.
            Args:
                num_pool: Total pools for running test
                num_cont: Total containers created on each pool
                total_count: Total times the test is run in a loop
                scm_size: SCM size used in the testing
                nvme_size: NVME size used in the testing
            Returns:
                None
        """
        pool = {}
        cont = {}

        for loop_count in range(0, total_count):
            self.log.info("Running test %s", loop_count)
            for val in range(0, num_pool):
                pool[val] = TestPool(self.context,
                                     dmg_command=self.get_dmg_command())
                pool[val].get_params(self)
                # Split total SCM and NVME size for creating multiple pools.
                temp = int(float(scm_size) / num_pool)
                pool[val].scm_size.update(str(temp))
                temp = int(float(nvme_size) / num_pool)
                pool[val].nvme_size.update(str(temp))
                pool[val].create()
                self.pool = pool[val]
                display_string = "pool{} space at the Beginning".format(val)
                self.pool.display_pool_daos_space(display_string)
                nvme_size_begin = self.pool.get_pool_free_space("NVME")
                for cont_val in range(0, num_cont):
                    cont[cont_val] = TestContainer(pool[val])

            for val in range(0, num_pool):
                display_string = "Pool{} space at the End".format(val)
                self.pool = pool[val]
                self.pool.display_pool_daos_space(display_string)
                nvme_size_end = self.pool.get_pool_free_space("NVME")
                pool[val].destroy()
                if nvme_size_begin != nvme_size_end:
                    self.fail("FAIL: NVME size not equal")

    def test_run(self, num_pool=1):
        """
        Method Descripton:
            This method is called with different test_cases.
            Args:
               num_pool: Total pools for running a test.
            Returns:
               None
        """
        num_jobs = self.params.get("no_parallel_job", '/run/ior/*')
        # Create a pool
        pool = {}

        # Iterate through IOR different ior test sequence
        for oclass, api, test, flags in product(self.ior_daos_oclass,
                                                self.ior_apis,
                                                self.ior_test_sequence,
                                                self.ior_flags):
            # Create the IOR threads
            threads = []
            for val in range(0, num_pool):
                pool[val] = TestPool(self.context,
                                     dmg_command=self.get_dmg_command())
                pool[val].get_params(self)
                # Split total SCM and NVME size for creating multiple pools.
                pool[val].scm_size.value = int(float(test[0])) / num_pool
                pool[val].nvme_size.value = int(float(test[1])) / num_pool
                pool[val].create()
                display_string = "pool{} space at the Beginning".format(val)
                self.pool = pool[val]
                self.pool.display_pool_daos_space(display_string)

                for thrd in range(0, num_jobs):
                    # Based on pools/jobs, split block size
                    if thrd == 0:
                        tmp = int(float(test[3])) / num_pool
                        test[3] = str(tmp)
                    # Add a thread for these IOR arguments
                    threads.append(threading.Thread(target=self.ior_thread,
                                                    kwargs={"pool": pool[val],
                                                            "oclass": oclass,
                                                            "api": api,
                                                            "test": test,
                                                            "flags": flags,
                                                            "results":
                                                            self.out_queue}))
            # Launch the IOR threads
            for thrd in threads:
                self.log.info("Thread : %s", thrd)
                thrd.start()
                time.sleep(5)
            # Wait to finish the threads
            for thrd in threads:
                thrd.join()

            # Verify the queue and make sure no FAIL for any IOR run
            # For 400G/800G, test should fail with ENOSPC.
            while not self.out_queue.empty():
                if self.out_queue.get() == "FAIL" and \
                   test[4] == "PASS":
                    self.fail("FAIL")

            for val in range(0, num_pool):
                display_string = "Pool{} space at the End".format(val)
                self.pool = pool[val]
                self.pool.display_pool_daos_space(display_string)
                self.pool.destroy()

    def test_nvme_pool_capacity(self):
        """Jira ID: DAOS-2085.

        Test Description:
            Purpose of this test is to verify whether DAOS stack
            report NOSPC when accessing data beyond pool size.
            Use Cases
            Test Case 1:
             1. Perform IO less than entire SSD disk space.
            Test Case 2:
             2. Perform IO beyond entire SSD disk space.
            Test Case 3:
             3. Create Pool/Container and destroy them several times.
            Test Case 4:
             4. Multiple Pool/Container testing.

        Use case:
        :avocado: tags=all,hw,large,nvme,full_regression
        :avocado: tags=nvme_pool_capacity
        """
        # Run test with one pool.
        self.log.info("Running Test Case 1 with one Pool")
        self.test_run(1)
        # Run test with two pools.
        self.log.info("Running Test Case 1 with two Pools")
        self.test_run(2)
        # Run Create/delete pool/container
        self.log.info("Running Test Case 3: Pool/Cont Create/Destroy")
        self.test_create_delete(10, 50, 100)