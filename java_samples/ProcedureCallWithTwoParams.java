public class ProcedureCallWithTwoParams {
   public static void main(String[] args) {
      myMethod("Hello", "World!");
   }

   static void myMethod(String arg1, String arg2) {
      System.out.println(arg1 + " " + arg2);
   }
}
