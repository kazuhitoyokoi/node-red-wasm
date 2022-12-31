package main
func flows() string { return `
[
    {
        "id": "8740d4c87aa226d5",
        "type": "http request",
        "z": "54d24c374bb820d2",
        "name": "",
        "method": "GET",
        "ret": "obj",
        "paytoqs": "ignore",
        "url": "https://api.wheretheiss.at/v1/satellites/25544",
        "tls": "",
        "persist": false,
        "proxy": "",
        "insecureHTTPParser": false,
        "authType": "",
        "senderr": false,
        "headers": [],
        "x": 330,
        "y": 120,
        "wires": [
            [
                "9ae96b81d3eb6fe2"
            ]
        ]
    },
    {
        "id": "7d5ef118b8c6f3cb",
        "type": "inject",
        "z": "54d24c374bb820d2",
        "name": "",
        "props": [
            {
                "p": "payload"
            },
            {
                "p": "topic",
                "vt": "str"
            }
        ],
        "repeat": "",
        "crontab": "",
        "once": true,
        "onceDelay": 0.1,
        "topic": "",
        "payload": "",
        "payloadType": "date",
        "x": 150,
        "y": 120,
        "wires": [
            [
                "8740d4c87aa226d5"
            ]
        ]
    },
    {
        "id": "9ae96b81d3eb6fe2",
        "type": "debug",
        "z": "54d24c374bb820d2",
        "name": "debug 1",
        "active": true,
        "tosidebar": true,
        "console": false,
        "tostatus": false,
        "complete": "false",
        "statusVal": "",
        "statusType": "auto",
        "x": 500,
        "y": 120,
        "wires": []
    }
]
`}
