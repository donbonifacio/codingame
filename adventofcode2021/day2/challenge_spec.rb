
require_relative './challenge.rb'

TEST_DATA = "forward 5
down 5
forward 8
up 3
down 8
forward 2".freeze

RSpec.describe Day2 do
  context 'when running test instructions' do
    it 'produces expected position factor' do
      expect(described_class.position_factor(TEST_DATA)).to eq(150)
    end
  end

  context 'when running input instructions' do
    it 'produces expected position factor' do
      INPUT_DATA = File.read('./input.txt')
      expect(described_class.position_factor(INPUT_DATA)).to eq(1_507_611)
    end
  end
end
