# json2csv

Converts a MailChimp streamed JSON to CSV format.

[![Build Status](https://travis-ci.org/hahow/json2csv.png?branch=master)](https://travis-ci.org/hahow/json2csv)

## Installation

If you have a working golang install, you can use `go get`.

```bash
go get github.com/hahow/json2csv
```

## Usage

```bash
usage: json2csv
    -k fields,to,output
    -i /path/to/input.json (optional; default is stdin)
    -o /path/to/output.csv (optional; default is stdout)
    --version
    -p print csv header row
    -h This help
```

To convert:

```json
["Email Address","First Name","Last Name","Address","Phone Number","MEMBER_RATING","OPTIN_TIME","OPTIN_IP","CONFIRM_TIME","CONFIRM_IP","LATITUDE","LONGITUDE","GMTOFF","DSTOFF","TIMEZONE","CC","REGION","LAST_CHANGED","LEID","EUID","NOTES"]
["appleseed_john@mac.com","John","Appleseed","","",2,"2018-06-12 07:53:28",null,"2018-06-12 07:53:28","76.216.210.0",null,null,null,null,null,null,null,"2018-06-12 07:53:28","3e4f9b94","fa7ae01bbebc",null]
```

to:

```csv
John,Appleseed,appleseed_john@mac.com
```

you would either

```bash
json2csv -k "First Name","Last Name","Email Address" -i input.json -o output.csv
```

or

```bash
cat input.json | json2csv -k "First Name","Last Name","Email Address" > output.csv
```
