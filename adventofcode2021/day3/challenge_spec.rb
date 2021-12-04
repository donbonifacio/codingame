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
INPUT_DATA = File.read('./input.txt').freeze

RSpec.describe Day3 do
  context 'when running test data' do
    it 'generates expected output' do
      expect(Day3.power_consumption(TEST_DATA)).to eq(198)
    end

    it 'generates expected gas output' do
      expect(Day3.gas_rating(TEST_DATA)).to eq(230)
    end
  end

  context 'when running input data' do
    it 'generates expected output' do
      expect(Day3.power_consumption(INPUT_DATA)).to eq(3_148_794)
    end

    it 'generates expected gas output' do
      expect(Day3.gas_rating(INPUT_DATA)).to eq(2795310)
    end
  end
end
