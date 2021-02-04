# ubuntu-mirror

Docker image to mirror ubuntu release and archive mirrors via rsync

## Environment Variables
| Name         | Default  | Required | Description                                                                                       |
|--------------|----------|----------|---------------------------------------------------------------------------------------------------|
| MIRROR_TYPE  | releases | no       | Type of mirror (releases or archive). Used for selecting the rsync script.                        |
| INTERVAL     | 360      | no       | How often the mirror should be synced in minutes (default 6 hours)                                |
| COUNTRY_CODE |          | no       | 2-letter country code to use a mirror in a different location                                     |
| RSYNC_SOURCE |          | no       | Custom rsync url to use; Rsync target must be the same type as MIRROR_TYPE; ignores COUNTRY_CODE; |

## Persistence
Mount the /data container path to a location on your host system

## Few Examples
Use default mirror (rsync://rsync.releases.ubuntu.com/ubuntu-releases/)
```bash
docker run -v "/hostPath:/data" twiese99/ubuntu-mirror
```
Use German mirror (rsync://de.rsync.releases.ubuntu.com/releases)
```bash
docker run -v "/hostPath:/data" --env COUNTRY_CODE='de' twiese99/ubuntu-mirror
```
Use archive mirror (rsync://archive.ubuntu.com/ubuntu/)
```bash
docker run -v "/hostPath:/data" --env MIRROR_TYPE='archive' twiese99/ubuntu-mirror
```
Use german archive mirror (rsync://de.rsync.archive.ubuntu.com/ubuntu)
```bash
docker run -v "/hostPath:/data" --env COUNTRY_CODE='de' --env MIRROR_TYPE='archive' twiese99/ubuntu-mirror
```
Use custom release mirror (rsync://ftp-stud.hs-esslingen.de/ubuntu.releases/)
```bash
docker run -v "/hostPath:/data" --env RSYNC_SOURCE='rsync://ftp-stud.hs-esslingen.de/ubuntu.releases/' twiese99/ubuntu-mirror
```