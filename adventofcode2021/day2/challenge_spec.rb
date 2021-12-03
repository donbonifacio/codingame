
require_relative './challenge.rb'

TEST_DATA = "forward 5
down 5
forward 8
up 3
down 8
forward 2".freeze

INPUT_DATA = File.read('./input.txt').freeze

RSpec.describe Day2 do
  context 'when running test instructions' do
    it 'produces expected position factor' do
      expect(Day2.position_factor(TEST_DATA)).to eq(150)
    end

    it 'produces expected aimed position factor' do
      expect(Day2Part2.position_factor(TEST_DATA)).to eq(900)
    end
  end

  context 'when running input instructions' do
    it 'produces expected position factor' do
      expect(Day2.position_factor(INPUT_DATA)).to eq(1_507_611)
    end

    it 'produces expected aimed position factor' do
      expect(Day2Part2.position_factor(INPUT_DATA)).to eq(1_880_593_125)
    end
  end
end
