# Throttle

A small client-server utility, that throttles commands sent to it from multiple
async sources. Also doubles as an error-notification handler for these jobs.

This is a rewrite of
[throttle-cli](https://git.sr.ht/~ferdinandyb/throttle-cli) in go, because
python startup times for the client was too slow for heavy usage.
