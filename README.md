# Usage

## Example

input.txt
```
"some string"
```

output.txt
```
some string
```

`$ ./substr input.txt \"(?P\<token\>.*?)\" {{.token}} > output.txt`