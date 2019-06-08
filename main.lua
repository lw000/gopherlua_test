function double(a)
    return a * 2
end

function maxmin( a, b )
	if a > b then
		return a, b
	else
		return b, a
	end
end

function max(a, b)
    if a > b then
        return a
    else
        return b
    end
end

min = function(a, b)
    if a > b then
        return b
    else
        return a
    end
end

function update(d)
    print(d)
    return 1
end

function main()
    -- print("lua call go func; " .. GoDouble(20))

    for i = 1, 10 do
        print(i .. i)
    end

    local ar = {a = 100, b = "bbbb", c = 10.49}
    for i, v in pairs(ar) do
        print(i ,v)
    end

    print("max: " .. max(100, 200))

    print("min: " .. min(100, 200))

    a, b = maxmin(100, 200)
    print("maxmin: " .. a .. ", " .. b)

    print(os.clock())

    print(os.date())

    -- print(levi.myfunc("11111"))

end

main()