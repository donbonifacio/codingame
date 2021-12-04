require_relative './challenge.rb'

TEST_DATA = "00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010".freeze
# INPUT_DATA = File.read('./input.txt').split(/\s/).map(&:to_i).freeze

RSpec.describe Day3 do
  context 'when running test data' do
    it 'generates expected output' do
      expect(Day3.power_consumption(TEST_DATA)).to eq(198)
    end
  end
end
