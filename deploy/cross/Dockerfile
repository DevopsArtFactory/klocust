# Copyright 2020 The klocust Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM dockercore/golang-cross:1.13.15 as base

ENV GO_VERSION 1.16.3

RUN rm -Rf /usr/local/go && mkdir /usr/local/go
RUN curl --fail --show-error --silent --location https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    | tar xz --directory=/usr/local/go --strip-components=1

# Cross compile klocust for Linux, Windows and MacOS
ARG GOOS
ARG GOARCH
ARG TAGS
ARG LDFLAGS

WORKDIR /klocust
COPY . ./

RUN if [ "$GOOS" = "darwin" ]; then export CC=o64-clang CXX=o64-clang++; fi; \
    GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=1 \
    go build -tags "${TAGS}" -ldflags "${LDFLAGS}" -o /build/klocust cmd/main.go
