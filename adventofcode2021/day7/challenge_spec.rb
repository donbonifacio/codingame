require_relative './challenge.rb'

TEST_DATA = '16,1,2,0,4,2,7,1,2,14'.freeze
INPUT_DATA = File.read('./input.txt').freeze

RSpec.describe Day7 do
  context 'when working with example data' do
    it 'has expected result for TEST_DATA' do
      expect(Day7.fuel_cost(TEST_DATA)).to eq(37)
    end

    it 'has expected result for TEST_DATA for expensive move' do
      expect(Day7.expensive_fuel_cost(TEST_DATA)).to eq(168)
    end
  end

  context 'when working with input data' do
    it 'has expected result for INPUT_DATA' do
      expect(Day7.fuel_cost(INPUT_DATA)).to eq(355_521)
    end

    it 'has expected result for TEST_DATA for expensive move' do
      expect(Day7.expensive_fuel_cost(INPUT_DATA)).to eq(100_148_777)
    end
  end
end
