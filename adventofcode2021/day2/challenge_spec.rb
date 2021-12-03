
require_relative './challenge.rb'

TEST_DATA = "forward 5
down 5
forward 8
up 3
down 8
forward 2".freeze

# TEST_DATA = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263].freeze
# INPUT_DATA = File.read('./input.txt').split(/\s/).map(&:to_i).freeze

RSpec.describe Day2 do
  context 'when running test instructions' do
    it 'produces expected position factor' do
      expect(described_class.position_factor(TEST_DATA)).to eq(150)
    end
  end
end
