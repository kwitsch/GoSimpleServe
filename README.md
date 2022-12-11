# GoSimpleServe

Tiny static http server container.  

All files in the `/static` folder will be served via http.  
The `/static/index.html` file is mandetory for the container startup.  

All environment variables are read at container start and won't refreshed afterwards.

## Configuration

| Environment variable | Type    | Default value | Description                                                                                                    |
|----------------------|---------|---------------|----------------------------------------------------------------------------------------------------------------|
| VERBOSE              | Boolean | false         | Additional information will be logged if enabled                                                               |
| ENDPOINT_FILES       | Boolean | false         | Enables the /files endpoint which returns a list of all files(including sub directories) in the /static folder |

## Config Template

If a `/config_template.yaml` file is present the `/config` endpoint will be enabled.  
This endpoint returns a json with values set in the container environment.  

### Structure

Each entry represents a field in the resulting json.  
Fields require these configurations:  

| Field        | Type                  | Description                                           |
|--------------|-----------------------|-------------------------------------------------------|
| envVariable  | String                | Environment variable name to map                      |
| defaultValue | String                | Default value which is used if environment is not set |
| variableType | Enum(bool,int,string) | variable type                                         |

### Example

Template:

```YAML
field-name1:
  envVariable: FIELD1
  defaultValue: false
  variableType: bool
field-name2:
  envVariable: FIELD2
  defaultValue: 1
  variableType: int
field-name3:
  envVariable: FIELD3
  defaultValue: test
  variableType: string
```

Witha an environment:

- FIELD1=true
- FIELD3=nothing

Results in an /config response of:

```JSON
{
    "field-name1": true,
    "field-name2": 1,
    "field-name3": "nothing"
}
```
