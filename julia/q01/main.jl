using DataStructures

function parseFile(path)
  lines = readlines(path)
  parts = split.(lines, "   ")
  left = (Int)[]
  right = (Int)[]
  for part in parts
    push!(left, parse(Int, part[1], base=10))
    push!(right, parse(Int, part[2], base=10))
  end

  return left, right
end

function part1()
  # path = joinpath(@__DIR__, "demo.txt")
  path = joinpath(@__DIR__, "main.txt")
  left, right = parseFile(path)
  sort!(left)
  sort!(right)
  diffs = abs.(left - right)
  res = sum(diffs)
  println("Answer: $res")
end

function part2()
  # path = joinpath(@__DIR__, "demo.txt")
  path = joinpath(@__DIR__, "main.txt")
  left, right = parseFile(path)
  rcnts = counter(right)
  sum = 0
  for el in left
    sum += get(rcnts, el, 0) * el
  end
  println("Answer: $sum")
end

# part1()
part2()
