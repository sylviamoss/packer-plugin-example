# Packer Plugin Example

This is a multi-plugin example built with real plugins. Even though this is a fully functional plugin,
this is not supposed to be used in real scenarios and should be used as a source of information of how 
multi-plugins are built.

To use the plugins from this example in real environments, please refer to the following: 
- Null Builder: https://www.packer.io/docs/builders/null
- Comment Provisioner: https://github.com/SwampDragons/packer-provisioner-comment .  
- Manifest Post-Processor: https://www.packer.io/docs/post-processors/manifest

## Building locally 

```
go build -o packer-plugin-example 
```

All multi-plugins are created as a new set. Sets have a `describe` command, which is used by 
Packer to get more information about all the possibilities inside a multi-plugin. 

```
./packer-plugin-example describe | jq .

{
  "version": "1.0.0-dev",
  "sdk_version": "1.6.6-dev",
  "builders": [
    "null"
  ],
  "post_processors": [
    "manifest"
  ],
  "provisioners": [
    "comment"
  ]
}

```