# Links

* Tutorial link
  https://macosxautomation.com/applescript/firsttutorial/18.html
* Key codes
  http://eastmanreference.com/complete-list-of-applescript-key-codes/
* Intro to extracting elements and properties
  http://macscripter.net/viewtopic.php?id=24754

## Good answers

* Accessibility inspector to view hierarychy of elements:
  https://apple.stackexchange.com/a/237110/2575630


# Script editor

* Watch log outputs in events-tab.


# Sample lines

```
tell application "Finder" to get the name of front Finder window
tell application "Finder" to close Finder window "Macintosh HD"
tell application "Finder" to get the index of Finder window "Macintosh HD"
tell application "Finder" to open the startup disk
tell application "Finder" to open home
tell application "Finder" to get the name of Finder Window 1
tell application "Finder" to  set the index of the last Finder window to 1
tell application "Finder" to get the target of the front window
tell application "Finder" to set the target of the last Finder window to home
tell application "Finder" to  set toolbar visible of the front Finder window to false
tell application "Finder" to set statusbar visible of  Finder window 1 to false
tell application "Finder" to set the sidebar width of  Finder window 1 to 240
tell application "Finder" to set the current view of the front Finder window to list view
 #otehr views -> column view/flow view/icon view/list view
tell application "Finder" to get the position of the front Finder window
tell application "Finder" to set the position of the front Finder window to {94, 134}
tell application "Finder" to get the bounds of the front window
tell application "Finder" to get the bounds of the window of the desktop
tell application "Safari" to open location "http://automator.us"
tell application "qutebrowser" to activate
```

# Concepts

## Properties of Finder Windows

* Name
* Index
* Target



## index

The value of this read-only property is a number corresponding to the window’s layer position in the stacking order of open Finder windows. On the computer, no two windows can occupy the same layer. One window is always on top of or in front of another window. The index property reflects this fact.

1 is on top

## special locations

* startup disk
* home

## index references

* by name:
    * Finder window "Documents"
* by numeric index:
    * Finder window 1
* by descriptive index:
    * the first Finder window
    * the second Finder window
    * the fifth Finder window
    * the 1st Finder window
    * the 23rd Finder window
* by relative position index:
    * the front Finder window
    * the middle Finder window
    * the back Finder window
    * the last Finder window
* by random index:
    * some Finder window

## verb used with windows

* get: used to access the current values of a window property
* set: used to apply a new value to a window property
* open: causes a window to display in the Finder
* close: causes a window to close
* select: causes a window to come to front

## Comments

### Single line comments

Begins with --

```
  -- This is a single line comment
```


# Collects

* Run from shell
    ```
    osascript -e 'display notification "Lorem ipsum dolor sit amet" with title "Title"'
    osascript -e 'tell app "System Events" to display dialog "Hello World"'
    osascript path/to/script.scpt
    ```
* Saving current app, switch and revert to the app
    ```
    set frontmostAppPath to (path to frontmost application) as text

    # Activate and work with another application.
    activate application "Reminders"
    delay 2

    # Make the previously frontmost (active) application frontmost again.
    activate application frontmostAppPath
    ```
* Send keystrokes to current app
    ```
    tell application "System Events" to tell (process 1 where frontmost is true)
        key code 119 using shift down
    end tell
    ```
* Get properties of an object
    ```
    tell application "Microsoft Outlook"
        get properties of front window
    end tell
    ```
* Get all running processes
    ```
    tell application "System Events"
        name of every process
    end tell
    ```
    * All windows (Not working)
        ```
        tell application "System Events"
            repeat with theProcess in processes
                if not background only of theProcess then
                    tell theProcess
                        set processName to name
                        name of every window
                    end tell
                end if
            end repeat
        end tell
        ```
* Get the names/classes of UI elements in a window
    ```
    tell application "System Events"
        class of UI elements of window "Accessibility" of process "System Preferences"
        name of UI elements of window "Accessibility" of process "System Preferences"
    end tell
    ```
* Get menu bar info
    Hierarchy: `Application/Process->menu bar 1->menu bar item "name"->menu "number"->menu item "name"->(then menu,menu item repeats to any depth)`
    Note that menu "" seems a repeat of previous item.
    ```
    activate application "Google Play Music Desktop Player"
    tell application "System Events"
        tell process "Google Play Music Desktop Player"
            set mbitems to menu bar items in menu bar 1
            repeat with mbit in mbitems
                set myname to name of mbit
                log "Menu bar item name is " & myname
                set mitems to menus in mbit
                repeat with each_menu in mitems
                    set mymname to name of each_menu
                    log "Menu name is " & mymname
                    set menu_items to menu items in each_menu
                    repeat with each_menu_item in menu_items
                        set menuitemname to name of each_menu_item
                        log "Menu item name is " & menuitemname
                    end repeat
                end repeat
            end repeat
        end tell
    end tell
    ```
* System preferences
    Apparently, system preferences has this hierarchy in Sierra: `pane->anchor`
    ```
    tell application "System Preferences"
        activate
        set current pane to pane "com.apple.preference.universalaccess"
        get the name of every anchor of pane id "com.apple.preference.universalaccess"
        reveal anchor "Mouse" of pane id "com.apple.preference.universalaccess"
    end tell
    ```
* Clicking a checkbox in system preferences
    ```
    tell application "System Preferences"
        activate
        set current pane to pane "com.apple.preference.universalaccess"
        reveal anchor "Mouse" of pane id "com.apple.preference.universalaccess"
    end tell
    tell application "System Events"
        tell application process "System Preferences"
            set frontmost to true
            click checkbox "Enable Mouse Keys" of window "Accessibility"
        end tell
    end tell
    ```
* Move mouse
    ```
    do shell script "/usr/local/bin/cliclick " & "c:23,23"
    ```


# My scripts

* Outlook move folder
    ```
    do shell script "/usr/local/bin/cliclick " & "c:23,23"
    ```
