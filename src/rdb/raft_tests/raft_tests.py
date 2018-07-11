#!/usr/bin/python
# Copyright (c) 2018 Intel Corporation
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
'''
Run the raft tests using make -C DIR tests, where DIR is the path to the raft
Makefile. Check the output for the number of "not ok" occurrences and return
this number as the return code.
'''
import subprocess
import sys
import os

TEST_NOT_RUN = -1

def number_of_failures():
    """
    Build and run the raft tests by calling the raft Makefile. Collect the
    output and examine for failures. If building or running fails, return
    TEST_NOT_RUN.
    """
    failures = 0
    successes = 0
    try:
        DIR = os.path.join(os.path.dirname(__file__), '../raft')
        res = subprocess.check_output(['make', '-C', DIR, 'tests'])
    except Exception as e:
        print("Building Raft Tests failed due to\n{}".format(e))
        return TEST_NOT_RUN

    for line in res.split('\n'):
        if line.startswith("not ok"):
            line = "FAIL: {}".format(line)
            failures += 1
        elif line.startswith("ok"):
            successes += 1
        print(line)

    if not successes and not failures:
        failures = TEST_NOT_RUN
    return failures

def main():
    failures = number_of_failures()
    if not failures:
        print("Raft Tests PASSED")
    elif failures == TEST_NOT_RUN:
        print("Raft Tests did not run")
    else:
        print("Raft Tests had {} failures".format(failures))
    sys.exit(failures)

if __name__ == "__main__":
    main()
