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
end
