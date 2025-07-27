#!/usr/bin/env python3
"""
Simple hello world script for testing Python execution
"""

import os
import sys
from datetime import datetime

def main():
    print("Hello from Cobra Script Center (Python)!")
    print(f"Current time: {datetime.now()}")
    print(f"Python version: {sys.version}")
    print(f"Working directory: {os.getcwd()}")
    
    # Check command line arguments
    if len(sys.argv) > 1:
        print("Arguments received:")
        for i, arg in enumerate(sys.argv[1:], 1):
            print(f"  {i}: {arg}")
    
    # Check environment variables
    name = os.environ.get('NAME')
    if name:
        print(f"Hello, {name}!")
    
    message = os.environ.get('MESSAGE')
    if message:
        print(f"Message: {message}")
    
    print("Python script execution completed successfully!")

if __name__ == "__main__":
    main()