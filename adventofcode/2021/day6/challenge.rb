# Solutions for Day 6.1
module Day6
  module_function

  def grow_fish(initial_state:, days:)
    fish = initial_state.split(',').map(&:strip).map(&:to_i)

    days.times do
      new_fish = 0
      fish = fish.map do |f|
        new_value = f - 1
        if new_value < 0
          new_value = 6
          new_fish += 1
        end
        new_value
      end
      fish += [8] * new_fish
    end

    { success: true, count: fish.count, state: fish.join(',') }
  end

  def optimized_grow_fish(initial_state:, days:)
    fish = load_fish(initial_state)
    days.times do
      fish, new_fish = handle_interval(fish)
      if new_fish > 0
        fish << { interval: 8, count: new_fish }
        fish = compress(fish)
      end
    end

    { success: true, count: fish.map { |f| f[:count] }.reduce(&:+) }
  end

  def load_fish(initial_state)
    initial_state
      .split(',')
      .map(&:strip)
      .map(&:to_i)
      .map { |interval| { interval: interval, count: 1 } }
  end

  def handle_interval(fish)
    new_fish = 0
    fish = fish.map do |data|
      data[:interval] -= 1
      if data[:interval] < 0
        data[:interval] = 6
        new_fish += data[:count]
      end
      data
    end
    [fish, new_fish]
  end

  def compress(fish)
    fish
      .group_by { |f| f[:interval] }
      .reduce([]) do |coll, entry|
        interval = entry.first
        fish = entry.last
        coll << { interval: interval, count: fish.map { |f| f[:count] }.reduce(&:+) }
      end
  end
end
