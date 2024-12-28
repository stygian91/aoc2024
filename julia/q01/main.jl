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

part1()
