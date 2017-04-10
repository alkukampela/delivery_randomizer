# Delivery humanizer

Input (payload as JSON)

```javascript
[{
    "key": "curl",
    "value": 9.0,
    "min": -10,
    "max": 10,
    "variance": 0.5
}, {
    "key": "line",
    "value": 90.0,
    "min": 0,
    "max": 360,
    "variance": 1.1
}]
```

Output

```javascript
[{
    "key": "curl",
    "value": 8.76
}, {
    "key": "line",
    "value": 90.3
}]
```
