import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day03 implements ISolver {

	@Override
	public void solvePart1() {
		try {
			File file = new File("./input/day3.txt");
			Scanner scanner = new Scanner(file);
			Pattern pattern = Pattern.compile("mul\\(\\d+,\\d+\\)", Pattern.CASE_INSENSITIVE);
			int sum = 0;
			while (scanner.hasNextLine()) {
				String line = scanner.nextLine();
				Matcher matcher = pattern.matcher(line);
				while (matcher.find()) {
					String found = matcher.group();
					String[] nums = found.substring(4, found.length() - 1).split(",");
					sum += Integer.parseInt(nums[0]) * Integer.parseInt(nums[1]);
				}
			}

			System.out.println(sum);
		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}

	private int getSum(String s) {
		String[] nums = s.substring(4, s.length() - 1).split(",");
		return Integer.parseInt(nums[0]) * Integer.parseInt(nums[1]);
	}

	@Override
	public void solvePart2() {
		try {
			File file = new File("./input/day3.txt");
			Scanner scanner = new Scanner(file);
			Pattern pattern = Pattern.compile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)", Pattern.CASE_INSENSITIVE);

			int sum = 0;
			boolean flag = true;

			while (scanner.hasNextLine()) {
				String line = scanner.nextLine();
				Matcher matcher = pattern.matcher(line);

				while (matcher.find()) {
					String found = matcher.group();
					if (found.equals("do()")) {
						flag = true;
					} else if (found.equals("don't()")){
					  flag = false;
					} else {
						if (flag) {
							sum += getSum(found);
						}
					}
				}
			}
			System.out.println("Final: "+ sum);

		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}

}
