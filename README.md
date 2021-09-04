# reader

Generate getters for golang using tags.

By default, getters for all private fields will be generated. To ignore, add `get:"-"` tag.


TODO
- change to getter
- allow embedded struct to be flattened, and optional prefix
- allow custom prefix, e.g. Get, or nested struct name
