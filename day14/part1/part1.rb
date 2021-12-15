# deez nuts

f = open("day14/input.txt")
if f == nil
  exit
end

lines = f.read.split("\n")

polymer_template = lines[0]
mappings = {}

lines[2..].map { |line| line.split(" -> ") }.each { |line|
  mappings[line[0]] = line[1]
}

(1..10).each {
  new_temp = ""
  (0..polymer_template.length - 2).each { |i|
    pair = polymer_template[i] + polymer_template[i + 1]
    mapping = mappings[pair]
    new_temp += polymer_template[i] + mapping
  }
  polymer_template = new_temp + polymer_template[-1]
}

quantities = {}

polymer_template.each_char.each { |c|
  quantities[c] ||= 0
  quantities[c] += 1
}

min = 10000000000
max = 0

quantities.each_value.each { |value|
  min = value unless min < value
  max = value unless max > value
}

puts "Solution: %d" % (max - min)