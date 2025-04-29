# Traefik Dyanmic Configuration Struct Definitions

This folder contains the struct definitions for the Traefik dynamic configuration. The structs are used to generate the dynamic configuration files for Traefik.

## Usage
To use the structs, you need to import the package and create an instance of the struct you want to use. Then, you can use the `json.Marshal` function to convert the struct to JSON format. (or vice-versa)


## License
This directory has been taken from https://github.com/traefik/traefik/tree/master/pkg/config/dynamic under the MIT license.

The MIT License (MIT)

Copyright (c) 2016-2020 Containous SAS; 2020-2025 Traefik Labs

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
