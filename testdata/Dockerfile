FROM rust:alpine

RUN \
    cargo install b3sum && \
    # Smoke test
    b3sum --version && \
    expect="718b749f12a61257438b2ea6643555fd995001c9d9ff84764f93f82610a780f2" && \
    actual="$(echo -n "This is a string" | b3sum --no-names -)" && \
    [ "$expect" = "$actual" ] || { echo "EXP:${expect}, ACT:${actual}";}

RUN \
    apk add --no-cache \
        jq

ENTRYPOINT [ "/usr/local/cargo/bin/b3sum" ]
CMD [ "--help" ]
