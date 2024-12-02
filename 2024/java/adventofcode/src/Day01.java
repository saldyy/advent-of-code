import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class Day01 implements ISolver {
	private Map<Integer, ArrayList<Integer>> processInput() {
		ArrayList<Integer> listA = new ArrayList<>();
		ArrayList<Integer> listB = new ArrayList<>();
		try {
			File file = new File("input/day1_1.txt");

			Scanner scanner = new Scanner(file);
			while (scanner.hasNext()) {
				int a = scanner.nextInt();
				listA.add(a);
				int b = scanner.nextInt();
				listB.add(b);
			}
			scanner.close();

		} catch (FileNotFoundException e) {
			throw new RuntimeException(e);
		}

		listA.sort((a, b) -> a - b);
		listB.sort((a, b) -> a - b);
		Map<Integer, ArrayList<Integer>> map = new HashMap<>();

		map.put(0, listA);
		map.put(1, listB);

		return map;
	}

	public void solvePart1() {
		Map<Integer, ArrayList<Integer>> map = processInput();
		ArrayList<Integer> listA = map.get(0);
		ArrayList<Integer> listB = map.get(1);

		System.out.println(listA.toString());
		System.out.println(listB.toString());
		int sum = 0;

		for (int i = 0; i < listA.size(); i++) {
			sum += Math.abs(listA.get(i) - listB.get(i));
		}

		System.out.println("Part 1:" + sum);
	}


	public void solvePart2() {
		Map<Integer, ArrayList<Integer>> map = processInput();
		ArrayList<Integer> listA = map.get(0);
		ArrayList<Integer> listB = map.get(1);

		Map<Integer, Integer> count = new HashMap<>();

		for (Integer item : listA) {
			count.put(item, 0);
		}

		for (Integer integer : listB) {
			if (count.containsKey(integer)) {
				count.put(integer, count.get(integer) + 1);
			}
		}
		int sum = 0;
		for (Map.Entry<Integer, Integer> entry: count.entrySet()) {
			int key = entry.getKey();
			int value = entry.getValue();
			sum += key * value;

		}

		System.out.println("Part 2:" + sum);
	}


}
