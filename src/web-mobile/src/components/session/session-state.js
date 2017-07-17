
const sessionState = {
    state: {
    "Id": 1,
    "AssignTo": {
      "id": 1,
      "username": "",
      "useridentity": ""
    },
    "expectingDate": "2017-07-17T00:00:00+08:00",
    "ExecutionDate": "0001-01-01T00:00:00Z",
    "originSession": {
      "id": 1,
      "name": "sessionName",
      "targetMuscle": "",
      "workouts": null,
      "repeat": 0,
      "weekly": "",
      "Plan": null
    },
    "Status": "assigned",
    "workouts": [
      {
        "Id": 1,
        "BelongSession": {
          "Id": 1,
          "AssignTo": null,
          "ExpectingDate": "0001-01-01T00:00:00Z",
          "ExecutionDate": "0001-01-01T00:00:00Z",
          "OriginSession": null,
          "Status": "",
          "Workouts": null
        },
        "MappedMovement": {
          "id": 1,
          "movement": {
            "id": 1,
            "targetMuscle": "胸",
            "secondaryMuscle": "肱三头肌",
            "name": "平板杠铃卧推",
            "description": "some description",
            "dividable": 0,
            "addBy": {
              "id": 1,
              "username": "",
              "useridentity": ""
            }
          },
          "sets": 4,
          "reps": 12,
          "sequence": 1,
          "needWarmup": 0,
          "Session": {
            "id": 1,
            "name": "",
            "targetMuscle": "",
            "workouts": null,
            "repeat": 0,
            "weekly": "",
            "Plan": null
          }
        },
        "Exercises": null,
        "Status": "0"
      }
    ]
  }
}