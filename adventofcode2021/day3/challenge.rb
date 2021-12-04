# Solutions for Day 3.1
module Day3
  module_function

  def power_consumption(report)
    report
      .lines
      .map(&:strip)
      .reduce({}, &method(:calculate_frequencies))
      .yield_self(&method(:calculate_gama))
      .yield_self(&method(:calculate_power_consumption))
  end

  def calculate_coll_frequencies(bag, lines)
    lines
      .map(&:strip)
      .reduce(bag, &method(:calculate_frequencies))
  end

  def calculate_frequencies(bag, line)
    line.each_char.with_index do |char, index|
      bag[index] ||= { '1' => 0, '0' => 0 }
      bag[index][char] += 1
    end
    bag
  end

  def calculate_gama(bag)
    bag
      .keys
      .sort
      .map do |position|
        if bag[position]['0'] > bag[position]['1']
          '0'
        else
          '1'
        end
      end.join
  end

  def calculate_power_consumption(gama)
    epsilon = invert(gama)
    gama.to_i(2) * epsilon.to_i(2)
  end

  def invert(gama)
    gama
      .chars
      .map { |char| char == '0' ? '1' : '0' }
      .join
  end

  #-----
  #----

  def gas_rating(report)
    lines = report.lines.map(&:strip)
    oxygen = calculate_gas_rating(lines, :most_frequent)
    co2 = calculate_gas_rating(lines, :less_frequent)

    oxygen.to_i(2) * co2.to_i(2)
  end

  def calculate_gas_rating(lines, filter_type)
    lines.first.length.times do |position|
      bag = calculate_coll_frequencies({}, lines).yield_self(&method(:enhance))
      data = bag[position]

      lines = keep_lines(lines, data, position, filter_type)
      # puts "--- #{position} - #{data}"
      # puts lines.inspect
      return lines.first if lines.length == 1
    end
    lines
  end

  def enhance(bag)
    bag.keys.reduce({}) do |new_bag, key|
      new_bag[key] = bag[key].clone
      if bag[key]['0'] > bag[key]['1']
        new_bag[key][:most_frequent] = '0'
        new_bag[key][:less_frequent] = '1'
      else
        new_bag[key][:most_frequent] = '1'
        new_bag[key][:less_frequent] = '0'
      end
      new_bag
    end
  end

  def keep_lines(lines, data, position, filter_type)
    lines.select do |line|
      line[position] == data[filter_type]
    end
  end
end
