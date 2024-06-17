# Go Windows Proxy

## Description

I have some bad experience using Windows 10 proxy settings in my old machine and want to make it easy to access either for win 10 or win 11

Current problem with proxy settings :

- Must search it or find it
- Must open settings -> Proxy settings
- Must click windows logo then search for "Proxy Settings"
- It's hard for non-technical users

Sometimes, I ask why it didn't included direcly in Network / Wi-Fi settings.

So I create my own...

It must be easy to access and to use.

## DEBUG

To try this version please build your own version or find the latest releases

`go build -o wproxy.exe`

Then you can open `wproxy.exe`

Expect to show the menu list hardcoded

Users can navigate

You can see your proxy status

I have not implement testing here
