# Go-Inca

---

### Features 💡

```
1. Set (key, val)
2. Get (key)
3. Set TTL(key, time)
4. Remove TTL(key)
5. Delete (key)
6. Set Eviction Policy (LRU/LFU)
    - LRU:
        - Capacity[MB]
    - LFU
        - Capacity[MB]
```

### Currently working on ⚙️

| Feature                                                      | Done(✅/❌) |
| ------------------------------------------------------------ | ----------- |
| Implement LRU                                                | ✅          |
| Implement LFU                                                | ❌          |
| Implement Set, Get and Delete for LRU                        | ✅          |
| Implement Set, Get and Delete for LFU                        | ❌          |
| Implement TTL                                                | ❌          |
| Optimize caching system                                      | ❌          |
| Add support for more data structures                         | ❌          |
| Make a network interface to communicated via TCP(Text Based) | ❌          |
| Make it work flawless with multiple concurrent users         | ❌          |
| Refactor DLL                                                 | ❌          |
| Refactor LRU                                                 | ❌          |

### Dev logs 🗒️

```
Tue 04 Jul 2023 12:47:42 PM IST
I have no idea about what I'm getting into, but yeah.
I'mma try my best to learn as much as I can, no matter what :)
```

```
Thu 06 Jul 2023 02:35:09 AM IST
Just implemented a doubly linkedlist and lru cache with basic functionality. Still hasn't found out the way to implement TTL.
This shit is fun af.
```

```
Thu 06 Jul 2023 01:43:52 PM IST
Everything thing is fast af, but I feel like I can improve it. Benchmarks are showing good results. I should do something to improve setting keys after capacity is over.
```
