# smoothbrain

SmoothBrain is designed to spread out nested json data for use-cases in NoSQL databases where data is ideally spread out thinly for easier query capabilities.

## Example

> Input

```json
{
    "name": "John Doe",
    "age": 36,
    "spouse": "Jane Doe",
    "children": [
        {
            "name": "Jimmy Doe",
            "age": 9
        },
        {
            "name": "Sally Doe",
            "age": 4
        }
    ],
    "vehicles": {
        "truck": {
            "make": "Ford"
        },
        "van": {
            "make": "Toyota"
        }
    }
}
```

> Output

```json
{
  "age": 36,
  "children": [
    {
      "age": 9,
      "name": "Jimmy Doe"
    },
    {
      "age": 4,
      "name": "Sally Doe"
    }
  ],
  "children.0": {
    "age": 9,
    "name": "Jimmy Doe"
  },
  "children.0.age": 9,
  "children.0.name": "Jimmy Doe",
  "children.1": {
    "age": 4,
    "name": "Sally Doe"
  },
  "children.1.age": 4,
  "children.1.name": "Sally Doe",
  "name": "John Doe",
  "spouse": "Jane Doe",
  "vehicles": {
    "truck": {
      "make": "Ford"
    },
    "van": {
      "make": "Toyota"
    }
  },
  "vehicles.truck": {
    "make": "Ford"
  },
  "vehicles.truck.make": "Ford",
  "vehicles.van": {
    "make": "Toyota"
  },
  "vehicles.van.make": "Toyota"
}
```

Not only does it flatten the data out but also has a record of every data structure in-between. If you had a DynamoDB table with a Hash Key of `PK` and a Sort Key of `SK` you can use each key in smoothbrain output as an individual item.

For example your query might be `PK = "PERSON#JohnDoe"` and `SK >= "children"`. That would give you the ability to get not only the flatten values of `children` but also the data in each nested section.
