# deez nuts

f = open("day14/input.txt")
if f == nil
  exit
end

lines = f.read.split("\n")

first_line = lines[0]
pairs = {}
first_line.split("").each_with_index { |c, i|
  if i != first_line.length - 1
    c = c + first_line[i + 1]
    pairs[c] = (pairs[c] || 0) + 1
  end
}

mappings = {}
lines[2..].map { |line| line.split(" -> ") }.each { |line|
  mappings[line[0]] = line[1]
}

(1..40).each {
  new_pairs = {}
  pairs.each_pair { |pair, times|
    mapping = mappings[pair]
    new_pair1 = pair[0] + mapping
    new_pair2 = mapping + pair[1]
    new_pairs[new_pair1] = (new_pairs[new_pair1] || 0) + times
    new_pairs[new_pair2] = (new_pairs[new_pair2] || 0) + times
  }
  pairs = new_pairs
}

quantities = {}

pairs.each_pair { |pair, times|
  quantities[pair[0]] = (quantities[pair[0]] || 0) + times / 2.0
  quantities[pair[1]] = (quantities[pair[1]] || 0) + times / 2.0
}

min = quantities.each_value.min.ceil
max = quantities.each_value.max.ceil

puts quantities
puts "Solution: %d" % (max - min)