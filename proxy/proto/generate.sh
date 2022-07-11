#!/bin/bash
#
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is used to generate the proxy proto APIs.

set -e

cd "$( dirname "${BASH_SOURCE[0]}" )"
mkdir -p reservation
protoc -I=.:../.. --go_out=./reservation --go_opt=paths=source_relative reservation.proto
mkdir -p httpovergrpc
protoc -I=.:../.. --go_out=./httpovergrpc --go_opt=paths=source_relative \
  --go-grpc_out=./httpovergrpc --go-grpc_opt=paths=source_relative  httpovergrpc.proto
