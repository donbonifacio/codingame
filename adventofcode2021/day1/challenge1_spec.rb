require_relative './challenge1.rb'

RSpec.describe Day1 do
  context 'when running #measure_increases' do
    it 'runs the default example' do
      data = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
      expect(Day1.measure_increases(data)).to eq(7)
    end

    it 'runs the input file' do
      data = File.read('./input.txt').split(/\s/).map(&:to_i)

      expect(data.count).to eq(2000)
      expect(Day1.measure_increases(data)).to eq(1195)
    end
  end
end
