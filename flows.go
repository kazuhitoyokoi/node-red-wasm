package main; func flows() string { return `
[
    {
        "id": "8740d4c87aa226d5",
        "type": "http request",
        "z": "9690b924941ce0a2",
        "name": "",
        "method": "GET",
        "ret": "txt",
        "paytoqs": "ignore",
        "url": "http://api.open-notify.org/iss-now.json",
        "tls": "",
        "persist": false,
        "proxy": "",
        "insecureHTTPParser": false,
        "authType": "",
        "senderr": false,
        "headers": [],
        "x": 490,
        "y": 200,
        "wires": [
            [
                "9ae96b81d3eb6fe2"
            ]
        ]
    },
    {
        "id": "7d5ef118b8c6f3cb",
        "type": "inject",
        "z": "9690b924941ce0a2",
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
        "repeat": "1",
        "crontab": "",
        "once": false,
        "onceDelay": 0.1,
        "topic": "",
        "payload": "",
        "payloadType": "date",
        "x": 310,
        "y": 200,
        "wires": [
            [
                "8740d4c87aa226d5"
            ]
        ]
    },
    {
        "id": "9ae96b81d3eb6fe2",
        "type": "debug",
        "z": "9690b924941ce0a2",
        "name": "debug 1",
        "active": true,
        "tosidebar": true,
        "console": false,
        "tostatus": false,
        "complete": "false",
        "statusVal": "",
        "statusType": "auto",
        "x": 660,
        "y": 200,
        "wires": []
    }
]`}
