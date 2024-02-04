#!/usr/bin/ruby
# frozen_string_literal: true

require 'json'

Node = Struct.new(:id, :children)
class Node
  def as_json(*)
    { id: id, children: children.map(&:as_json) }
  end
end

def main
  node_children = (2..5)

  n0 = Node.new 0, []
  to_fill = [n0]
  total_nodes = 1
  max_nodes = 200
  # ids = (1..10_000).to_a
  ids = (1..4000).to_a + (8000..12_000).to_a

  100.times do
    ids.shuffle!
    ids.shift(10)
  end
  ids.shuffle!

  while total_nodes < max_nodes
    current_node = to_fill.shift
    node_children.to_a.sample.times do
      n = Node.new ids.shift, []
      current_node.children.append n
      total_nodes += 1
      to_fill.append n
    end
  end

  case ARGV[0]
  when 'pretty'
    puts JSON.pretty_generate n0.as_json
  when 'irb'
    binding.irb
  else
    puts JSON.generate n0.as_json
  end
end

main
