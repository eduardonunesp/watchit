# WatchIt!
Watch a file change and run a command instantly

## How to use
Watchit is very simple to use, you just need to inform the watchable and the command which will run after any change in the watchable, for instance:

```
watchit config.txt reload_some_app
```

The line above mean: For any update in file `config.txt` execute the command `relaod_some_app`, very simple, isn't ?
You can watch an entire directory as well, just inform the directory path.

```
watchit config_dir reload_some_app
````

Same command from above, but now it will watch the entire directory tree from the root `config_dir`

### Notes

* The project still is on his infancy, contributions are welcome.
* License MIT
