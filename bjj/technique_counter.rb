# Sums techniques
module TechniqueCalculator
  module_function

  def run(data)
    parse(data)
  end

  def parse(data)
    data
      .lines
      .map(&:strip)
      .reject(&:empty?)
      .yield_self(&method(:explode))
      .yield_self(&method(:resolve_multiplier))
      .yield_self(&method(:count))
  end

  def explode(data)
    data.reduce([]) do |coll, line|
      coll + line.split(/,/).map(&:strip)
    end
  end

  def resolve_multiplier(data)
    data.map do |line|
      multiplier = line.match(/(\d)x\s+/)
      if multiplier
        raw = multiplier[0].strip
        technique = line.gsub!(raw, '').strip
        [technique] * raw.delete('x').strip.to_i
      else
        line
      end
    end.flatten
  end

  def count(data)
    data.reduce({}) do |bag, line|
      bag[line] ||= 0
      bag[line] += 1
      bag
    end.sort_by { |_, v| -v }
  end
end

# puts '-- Submissions'
# TechniqueCalculator.run(File.read('./submissions.txt'))

# puts '-- Taps'
# TechniqueCalculator.run(File.read('./taps.txt'))
