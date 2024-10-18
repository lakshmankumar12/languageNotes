# links

https://testautomationu.applitools.com/robot-framework-refresh/chapter1.html
* Current at ch-4


# Notes

* Single-space separation means same token
  * two or more spaces - separates key word names, args (also from other args)

# comments

```
# line beginning with # is a comment
```


# command line

* https://baishanlu.gitbooks.io/robot-framework-cn-en-user-manual/content/6appendices/62all_command_line_options.html

```
-T  timestamp output
-d dir  alls results in the dir

```

# constructs

## variables

```robot
## basic
###   note the {..} is mandatory
${SCALAR}
@{LIST}
&{DICT}
%{ENV_VAR}

## multi-line variable
${example_regex}=  catenate  SEPARATOR=
...  (?m)Setting IP address to [0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\n
...  Setting MAC address to [0-9a-f]{2}:[0-9a-f]{2}:[0-9a-f]{2}:[0-9a-f]{2}:[0-9a-f]{2}:[0-9a-f]{2}\\n
...  Setting IP forwarding kernel options


## simple set
${hi} =    Set Variable    Hello, world!
```


## loops

```robot




FOR   ${i}  IN RANGE  15
   ....
   Exit For Loop
END

FOR    ${key}    IN    @{mydict.keys()}
    Log    ${mydict["${key}"]}
END

```

# dictionay

```robot

&{args}=    Create Dictionary    serials=${devices_serial_no_list}

## search: add dictionary
Set To Dictionary    ${sas_curlParams}    cert=${sasemulator_crt}   key=${sasemulator_key}

&{params}=    Copy Dictionary    ${onyxedge_orc8r_dict}

```



# Popular Keywords

```robot

# check if dict has a key
Run keyword if  'key1' in $TestCase  Input Text  ...

${cell_Id}=     Evaluate    ${cell_Id} + 1

${file_exists}=  Run Keyword and Return Status    SSHLibrary.File Should Exist    ${fname}

##run a should and return its value
${status}=    Run Keyword And Return Status   Should Contain   ${out}   Server started at http://localhost:8380
IF    ${status} == False
...
END

## json .. back and forth
${robo_dict}=        evaluate        json.loads('''${json_string}''')    json
set to dictionary    ${robot_dict["vt"]}    dp=the new value
${json_string}=      evaluate        json.dumps(${robo_dict})                 json

```


# libraries

## ssh library

http://robotframework.org/SSHLibrary/SSHLibrary.html

* Open Connection
* Login
* Switch Connection
* Put File
* Get File


## sleep for some time

```
Sleep   20seconds

```

