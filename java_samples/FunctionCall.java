public class FunctionCall {
   public static void main(String[] args) {
      System.out.println("Hello from main!");
      System.out.println(myMethod());
   }

   static String myMethod() {
      return "Hello from myMethod!";
   }
}
