
# Solutions for Day 1.1
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
end
