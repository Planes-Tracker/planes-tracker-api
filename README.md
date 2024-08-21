# Planes tracker API

This project is still under development. Expect changes and bugs.

See the [changelog](/CHANGELOG.md) for the latest updates.

## Table of content

-   [**Installation**](#installation)
-   [**Compiling from source**](#compiling-from-source)
-   [**Configuration**](#configuration)
-   [**Config details**](#config-details)
-   [**Copyright**](#copyright)

## Installation

-   Download [go](https://go.dev/dl/) (go 1.20 required).
-   Download or clone the project.
-   Download the binary from the [Releases](../../releases) or [build it](#compiling-from-source) yourself.
-   [Configure the app](#configuration).

## Compiling from source

Use `go build` in the project directory.

## Configuration

The configuration details must be set inside a `.env` file at the root of the project. An exemple is provided inside [`.env.example`](/.env.example).

## Config details

| Item                         | Values                  | Meaning                                                           |
| ---------------------------- | ----------------------- | ----------------------------------------------------------------- |
| `TRACKER_LOCATION_LATITUDE`  | `floating point number` | Latitude of the center of the area to be covered                  |
| `TRACKER_LOCATION_LONGITUDE` | `floating point number` | Longitude of the center of the area to be covered                 |
| `TRACKER_RADIUS_DISTANCE`    | `number`                | Radius distance from the center of the area to be covered (in Km) |

## Copyright

See the [license](/LICENSE).
