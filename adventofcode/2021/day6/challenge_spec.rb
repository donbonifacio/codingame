require_relative './challenge.rb'

INPUT_DATA = File.read('./input.txt').freeze

RSpec.describe Day6 do
  context 'considering basic scenarios' do
    [
      {
        scenario: { initial_state: '6', days: 1 },
        expected: { success: true, count: 1, state: '5' }
      },
      {
        scenario: { initial_state: '0', days: 1 },
        expected: { success: true, count: 2, state: '6,8' }
      },
      {
        scenario: { initial_state: '3,4,3,1,2', days: 1 },
        expected: { success: true, count: 5, state: '2,3,2,0,1' }
      },
      {
        scenario: { initial_state: '3,4,3,1,2', days: 2 },
        expected: { success: true, count: 6, state: '1,2,1,6,0,8' }
      },
      {
        scenario: { initial_state: '3,4,3,1,2', days: 18 },
        expected: { success: true, count: 26, state: '6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8' }
      }
    ].each do |test_data|
      it "grows as expected after one day when starting for #{test_data}" do
        result = Day6.grow_fish(test_data[:scenario])
        expect(result).to eq(test_data[:expected])

        result2 = Day6.optimized_grow_fish(test_data[:scenario])
        expect(result2[:count]).to eql(result[:count])
      end
    end

    it 'generates passing result for 256 days' do
      result = Day6.optimized_grow_fish(initial_state: '3,4,3,1,2', days: 256)
      expect(result[:count]).to eq(26_984_457_539)
    end
  end

  context 'when running input data' do
    it 'generates passing result for 80 days' do
      result = Day6.grow_fish(initial_state: INPUT_DATA, days: 80)
      expect(result[:count]).to eq(386_536)
    end

    it 'generates passing result for 256 days' do
      result = Day6.optimized_grow_fish(initial_state: INPUT_DATA, days: 256)
      expect(
        result[:count]).to eq(1_732_821_262_171)
    end
  end
end
