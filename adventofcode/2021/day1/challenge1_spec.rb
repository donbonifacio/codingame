require_relative './challenge1.rb'

TEST_DATA = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263].freeze
INPUT_DATA = File.read('./input.txt').split(/\s/).map(&:to_i).freeze

RSpec.describe Day1 do
  context 'when working with provided data' do
    it 'has expected data for TEST_DATA' do
      expect(TEST_DATA.count).to eq(10)
    end

    it 'has expected data for INPUT_DATA' do
      expect(INPUT_DATA.count).to eq(2000)
    end
  end

  context 'when running #measure_increases' do
    it 'runs the default example' do
      expect(Day1.measure_increases(TEST_DATA)).to eq(7)
    end

    it 'runs the input file' do
      expect(INPUT_DATA.count).to eq(2000)
      expect(Day1.measure_increases(INPUT_DATA)).to eq(1195)
    end
  end

  context 'when running #rolling_measure_increases' do
    it 'runs the default example' do
      expect(Day1.rolling_measure_increases(TEST_DATA)).to eq(5)
    end

    it 'runs the input file' do
      expect(Day1.rolling_measure_increases(INPUT_DATA)).to eq(1235)
    end
  end
end
