require_relative './technique_counter.rb'

RSpec.describe TechniqueCalculator do
  context 'when dealing with basic data' do
    it 'it counts direct techniques on single line' do
      expect(described_class.run('ArmBar, Kimura'))
        .to eq([['ArmBar', 1], ['Kimura', 1]])
    end

    it 'it counts direct techniques on single line ignoring white spaces' do
      expect(described_class.run('ArmBar,Kimura'))
        .to eq([['ArmBar', 1], ['Kimura', 1]])
    end

    it 'it counts techniques on multiple lines' do
      expect(described_class.run("ArmBar, Kimura\nOmoplata"))
        .to eq([['ArmBar', 1], ['Kimura', 1], ['Omoplata', 1]])
    end

    it 'it counts the same techniques on multiple lines' do
      expect(described_class.run("ArmBar, ArmBar\nArmBar"))
        .to eq([['ArmBar', 3]])
    end

    it 'counts techniques with multiplier' do
      expect(described_class.run('ArmBar, 3x Kimura'))
        .to eq([['Kimura', 3], ['ArmBar', 1]])
    end
  end
end
