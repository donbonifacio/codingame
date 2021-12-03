require_relative './challenge.rb'

# TEST_DATA = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263].freeze
# INPUT_DATA = File.read('./input.txt').split(/\s/).map(&:to_i).freeze

RSpec.describe Day1 do
  context 'when working with provided data' do
    it 'has expected data for TEST_DATA' do
      expect(TEST_DATA.count).to eq(10)
    end

    xit 'has expected data for INPUT_DATA' do
      expect(INPUT_DATA.count).to eq(2000)
    end
  end
end
