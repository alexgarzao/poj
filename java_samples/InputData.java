public class InputData {
    public static String name;
    public static int age;

    public static void main(String[] args) {
        // Using Console to input data from user
        name = System.console().readLine();
        age = Integer.parseInt(System.console().readLine());

        System.out.println("You entered string " + name);
    }
}
