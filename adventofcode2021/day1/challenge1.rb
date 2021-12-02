
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
    return 0 if previous > current

    1
  end
end
