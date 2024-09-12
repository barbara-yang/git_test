wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
local requests = 1

function request()
    requests = math.random(1,10000000)
    body = string.format("username=test%d&password=test%d",requests,requests)
    return wrk.format('POST', nil, nil, body)
end
function response(status, headers, body)
    if (status ~= 302)
    then
        print(status)
    end
end

