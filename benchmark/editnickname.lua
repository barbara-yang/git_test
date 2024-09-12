wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

function request()
    requests = math.random(1,10000000)
    wrk.headers["Cookie"] ="session=test-session;userid="..requests
    body = string.format("nickname=test%d",requests)
    return wrk.format('POST', nil, nil, body)
end
function response(status, headers, body)
    if (status ~= 302)
    then
        print(status)
    end
end