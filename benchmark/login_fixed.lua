wrk.method = "POST"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

-- 预定义用户账号
local users = {}
for i = 1, 200 do
    table.insert(users, {username = "test" .. i, password = "test" .. i})
end

-- 选择用户的索引
local counter = 1

function request()
    local user_index = (counter % 200) + 1
    local user = users[user_index]
    counter = counter + 1
    local body = string.format("username=%s&password=%s", user.username, user.password)
    return wrk.format('POST', nil, nil, body)
end

function response(status, headers, body)
    if (status ~= 302) then
        print(status)
    end
end


