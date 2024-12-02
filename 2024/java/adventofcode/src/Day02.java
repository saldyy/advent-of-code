import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class Day02 implements ISolver {
	private ArrayList<int[]> input ;
	Day02() {
		this.input = new ArrayList<int[]>();
	}

	void processInput() {
		try {
			this.input = new ArrayList<int[]>();
			File file = new File("input/day2.txt");
			Scanner scanner = new Scanner(file);
			while (scanner.hasNext()) {
				String line = scanner.nextLine();
				String[] split = line.split(" ");
				int[] array = new int[split.length];

				int i = 0;
				for (String s: split) {
					array[i] = Integer.parseInt(s);
					i++;
				}
				this.input.add(array);
			}

			scanner.close();
		} catch (FileNotFoundException e) {
			System.out.println("File not found");
		}
	}

	private boolean checkIsSafeReport(int[] array) {
		boolean isSafeReport = true;
		int direction = array[0] - array[1];

		for (int j=1; j<array.length; j++) {
			int rawDiff = array[j-1] - array[j];
			if (direction * rawDiff < 0) {
				isSafeReport = false;
				break;
			}
			int diff = Math.abs(rawDiff);
			if (diff< 1 || diff>3) {
				isSafeReport = false;
				break;
			}
		}

		return isSafeReport;
	}

	public void solvePart1() {
		this.processInput();
		int result = 0;

		for (int i = 0; i < this.input.size(); i++) {
			int[] array = this.input.get(i);
			boolean isSafeReport = checkIsSafeReport(array);
			if (isSafeReport) {
				result++;
			}
		}

		System.out.printf("Result: %d\n", result);
	}


	public void solvePart2() {
		this.processInput();
		int result = 0;

		for (int i = 0; i < this.input.size(); i++) {
			int[] array = this.input.get(i);
			boolean isSafeReport = this.checkIsSafeReport(array);

			if (!isSafeReport) {
				for (int j=0; j<array.length; j++) {
					int[] clonedArray = new int[array.length-1];
					int index = 0;

					for (int k=0; k<array.length; k++) {
						if (k==j)	 {
							continue;
						}
						clonedArray[index++] = array[k];
					}

//					System.out.println(Arrays.toString(clonedArray));
					isSafeReport = checkIsSafeReport(clonedArray);
					if (isSafeReport) {
						break;
					}

				}

			}

			if (isSafeReport) {
				result++;
			}
		}

		System.out.printf("Result: %d\n", result);
	}
}
