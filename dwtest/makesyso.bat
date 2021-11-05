@echo off
windres -i dwtest.rc -o dwtest_windows_amd64.syso --input-format=rc --output-format=coff --target=pe-x86-64 -DDW64
windres -i dwtest.rc -o dwtest_windows_386.syso --input-format=rc --output-format=coff --target=pe-i386
