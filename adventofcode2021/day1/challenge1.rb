
# Solutions for Day 1.1
module Day1
  def self.measure_increases(measures)
    data = measures.reduce(count: 0, previous: nil) do |acc, measure|
      previous = acc[:previous]
      count = acc[:count] + add_increases(previous, measure)
      # puts "#{previous} #{measure} => #{count}"

      { count: count, previous: measure }
    end
    data[:count]
  end

  def self.add_increases(previous, current)
    return 0 if previous.nil?
    return 0 if previous >= current

    1
  end

  def self.rolling_measure_increases(measures)
    index = 0
    windows = []

    loop do
      current_slice = measures.slice(index, 3)
      break if current_slice.length != 3
      windows << current_slice
      index += 1
    end

    sums = windows.map { |window| window.reduce(&:+) }
    measure_increases(sums)
  end
end
