#!/usr/bin/env python3
# -*- coding: utf-8 -*-


import ctypes

lib = ctypes.CDLL('test.so')
lib.test(ctypes.c_double(5))